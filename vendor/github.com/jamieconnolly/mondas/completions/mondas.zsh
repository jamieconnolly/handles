#compdef mondas

local args cmds line ret=1 state

_arguments -C \
  '--help[Display help information]' \
  '--version[Display version information]' \
  '1: :->cmds' \
  '*: :->args' && ret=0

case $state in
  args)
    args=(${(f)"$(_call_program arguments ${service} completions ${line[@]:0:$((CURRENT-1))} 2>/dev/null)"})
    [[ $#args -gt 0 ]] && _values "${service} arguments" "${args[@]}" && ret=0
    ;;
  cmds)
    cmds=(${(f)"$(_call_program commands ${service} completions 2>/dev/null)"})
    _describe -t commands "${service} commands" cmds && ret=0
    ;;
esac

return ret
