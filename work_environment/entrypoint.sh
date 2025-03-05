gitlab-runner register \
    --non-interactive \
    --url http://gitlab \
    --token glrt-t3_12jk6E5ELrNeYptSXfEu \
    --executor "shell" \
    --description "shell-runner"

gitlab-runner run
