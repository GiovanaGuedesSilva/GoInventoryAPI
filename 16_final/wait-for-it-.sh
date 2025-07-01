#!/usr/bin/env bash
# Esse script verifica se um host e porta TCP estão disponíveis
# Ele é usado para aguardar um serviço (como o MySQL) estar pronto antes de iniciar a aplicação

# Armazena o nome do script
WAITFORIT_cmdname=${0##*/}

# Função auxiliar para imprimir mensagens de erro (caso o modo silencioso não esteja ativado)
echoerr() { if [[ $WAITFORIT_QUIET -ne 1 ]]; then echo "$@" 1>&2; fi }

# Função de uso/ajuda
usage()
{
    cat << USAGE >&2
Usage:
    $WAITFORIT_cmdname host:port [-s] [-t timeout] [-- command args]
    -h HOST | --host=HOST       Host/IP para testar
    -p PORT | --port=PORT       Porta TCP para testar
                                Alternativamente, você pode passar host:porta direto
    -s | --strict               Executa o comando somente se o teste for bem-sucedido
    -q | --quiet                Silencia as mensagens de status
    -t TIMEOUT | --timeout=TIMEOUT
                                Tempo limite (em segundos) para aguardar
    -- COMMAND ARGS             Comando a executar depois que o host:porta estiverem disponíveis
USAGE
    exit 1
}

# Função que espera pela disponibilidade do host:porta
wait_for()
{
    if [[ $WAITFORIT_TIMEOUT -gt 0 ]]; then
        echoerr "$WAITFORIT_cmdname: aguardando $WAITFORIT_TIMEOUT segundos por $WAITFORIT_HOST:$WAITFORIT_PORT"
    else
        echoerr "$WAITFORIT_cmdname: aguardando indefinidamente por $WAITFORIT_HOST:$WAITFORIT_PORT"
    fi

    WAITFORIT_start_ts=$(date +%s)
    while :
    do
        # Tenta conectar via netcat se estiver disponível
        if [[ $WAITFORIT_ISBUSY -eq 1 ]]; then
            nc -z $WAITFORIT_HOST $WAITFORIT_PORT
            WAITFORIT_result=$?
        else
            # Alternativa usando redirecionamento TCP (funciona em Bash puro)
            (echo -n > /dev/tcp/$WAITFORIT_HOST/$WAITFORIT_PORT) >/dev/null 2>&1
            WAITFORIT_result=$?
        fi

        # Se conseguiu conectar, sai do loop
        if [[ $WAITFORIT_result -eq 0 ]]; then
            WAITFORIT_end_ts=$(date +%s)
            echoerr "$WAITFORIT_cmdname: $WAITFORIT_HOST:$WAITFORIT_PORT disponível após $((WAITFORIT_end_ts - WAITFORIT_start_ts)) segundos"
            break
        fi

        # Aguarda 1 segundo antes de tentar novamente
        sleep 1
    done

    return $WAITFORIT_result
}

# Wrapper para suportar timeout com sinal de interrupção
wait_for_wrapper()
{
    if [[ $WAITFORIT_QUIET -eq 1 ]]; then
        timeout $WAITFORIT_BUSYTIMEFLAG $WAITFORIT_TIMEOUT $0 --quiet --child --host=$WAITFORIT_HOST --port=$WAITFORIT_PORT --timeout=$WAITFORIT_TIMEOUT &
    else
        timeout $WAITFORIT_BUSYTIMEFLAG $WAITFORIT_TIMEOUT $0 --child --host=$WAITFORIT_HOST --port=$WAITFORIT_PORT --timeout=$WAITFORIT_TIMEOUT &
    fi

    WAITFORIT_PID=$!
    trap "kill -INT -$WAITFORIT_PID" INT
    wait $WAITFORIT_PID
    WAITFORIT_RESULT=$?

    if [[ $WAITFORIT_RESULT -ne 0 ]]; then
        echoerr "$WAITFORIT_cmdname: tempo esgotado após $WAITFORIT_TIMEOUT segundos para $WAITFORIT_HOST:$WAITFORIT_PORT"
    fi

    return $WAITFORIT_RESULT
}

# Processamento dos argumentos da linha de comando
while [[ $# -gt 0 ]]
do
    case "$1" in
        *:* )
        # Suporta o formato host:porta direto
        WAITFORIT_hostport=(${1//:/ })
        WAITFORIT_HOST=${WAITFORIT_hostport[0]}
        WAITFORIT_PORT=${WAITFORIT_hostport[1]}
        shift 1
        ;;
        --child)
        WAITFORIT_CHILD=1
        shift 1
        ;;
        -q | --quiet)
        WAITFORIT_QUIET=1
        shift 1
        ;;
        -s | --strict)
        WAITFORIT_STRICT=1
        shift 1
        ;;
        -h)
        WAITFORIT_HOST="$2"
        shift 2
        ;;
        --host=*)
        WAITFORIT_HOST="${1#*=}"
        shift 1
        ;;
        -p)
        WAITFORIT_PORT="$2"
        shift 2
        ;;
        --port=*)
        WAITFORIT_PORT="${1#*=}"
        shift 1
        ;;
        -t)
        WAITFORIT_TIMEOUT="$2"
        shift 2
        ;;
        --timeout=*)
        WAITFORIT_TIMEOUT="${1#*=}"
        shift 1
        ;;
        --)
        shift
        WAITFORIT_CLI=("$@")
        break
        ;;
        --help)
        usage
        ;;
        *)
        echoerr "Argumento desconhecido: $1"
        usage
        ;;
    esac
done

# Verificação de obrigatoriedade dos parâmetros
if [[ "$WAITFORIT_HOST" == "" || "$WAITFORIT_PORT" == "" ]]; then
    echoerr "Erro: host e porta são obrigatórios."
    usage
fi

# Define valores padrão
WAITFORIT_TIMEOUT=${WAITFORIT_TIMEOUT:-15}
WAITFORIT_STRICT=${WAITFORIT_STRICT:-0}
WAITFORIT_CHILD=${WAITFORIT_CHILD:-0}
WAITFORIT_QUIET=${WAITFORIT_QUIET:-0}

# Verifica se o timeout disponível é da versão BusyBox
WAITFORIT_TIMEOUT_PATH=$(type -p timeout)
WAITFORIT_TIMEOUT_PATH=$(realpath $WAITFORIT_TIMEOUT_PATH 2>/dev/null || readlink -f $WAITFORIT_TIMEOUT_PATH)
WAITFORIT_BUSYTIMEFLAG=""

if [[ $WAITFORIT_TIMEOUT_PATH =~ "busybox" ]]; then
    WAITFORIT_ISBUSY=1
    if timeout &>/dev/stdout | grep -q -e '-t '; then
        WAITFORIT_BUSYTIMEFLAG="-t"
    fi
else
    WAITFORIT_ISBUSY=0
fi

# Execução do script principal
if [[ $WAITFORIT_CHILD -gt 0 ]]; then
    wait_for
    WAITFORIT_RESULT=$?
    exit $WAITFORIT_RESULT
else
    if [[ $WAITFORIT_TIMEOUT -gt 0 ]]; then
        wait_for_wrapper
        WAITFORIT_RESULT=$?
    else
        wait_for
        WAITFORIT_RESULT=$?
    fi
fi

# Executa o comando final (se fornecido)
if [[ $WAITFORIT_CLI != "" ]]; then
    if [[ $WAITFORIT_RESULT -ne 0 && $WAITFORIT_STRICT -eq 1 ]]; then
        echoerr "$WAITFORIT_cmdname: modo estrito ativado, não executando o subprocesso"
        exit $WAITFORIT_RESULT
    fi
    exec "${WAITFORIT_CLI[@]}"
else
    exit $WAITFORIT_RESULT
fi
