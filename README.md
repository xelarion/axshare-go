Web: https://github.com/XanderCheung/axshare

Init env software

    chmod +x scripts/init_env.sh
    ./scripts/init_env.sh

Dockerize axshare-vue

    // build app
    docker build -t ervincheung/axshare-go .
    // run single app
    docker run -d -p 10524:10524 --rm --name dockerize-axshare_go ervincheung/axshare-go
***
