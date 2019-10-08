# config valid for current version and patch releases of Capistrano
lock '~> 3.11.0'

# 项目名称
set :application, 'axshare-go'
# 仓库地址
set :repo_url, 'git@github.com:ErvinCheung/axshare-go.git'

# Default branch is :master
# ask :branch, `git rev-parse --abbrev-ref HEAD`.chomp

# Default deploy_to directory is /var/www/my_app_name
# set :deploy_to, "/var/www/my_app_name"
# 需要部署到服务器的位置
set :deploy_to, '~/web/axshare-go'

# bundle 报错
# set :bundle_flags, '--deployment --quiet --full-index'

# Default value for :format is :airbrussh.
# set :format, :airbrussh

# You can configure the Airbrussh format using :format_options.
# These are the defaults.
# set :format_options, command_output: true, log_file: "log/capistrano.log", color: :auto, truncate: :auto

# Default value for :pty is false
# set :pty, true

# Default value for :linked_files is []
# append :linked_files, "config/database.yml"
# 去掉注释，并加上 "config/master.key", 如果有storage.yml 也加上
append :linked_files, "config/config.yaml", ".env", ".axshare.env", ".mysql.env"

append :linked_dirs, "log"

# Default value for default_env is {}
# set :default_env, { path: "/opt/ruby/bin:$PATH" }ne

# Default value for local_user is ENV['USER']
# set :local_user, -> { `git config user.name`.chomp }

# Default value for keep_releases is 5
# set :keep_releases, 5
set :keep_releases, 3

# Uncomment the following to require manually verifying the host key before first deploy.
# set :ssh_options, verify_host_key: :secure

namespace :deploy do
  before :starting, :ensure_user do
    run_locally do
        # execute "sudo apt-get update"
    end
  end
end
