set :application, "go_demo"

set :repository, "git@github.com:adeven/go_demo.git"

ssh_options[:forward_agent] = true

set :user, 'go_demo'
set :deploy_to, "/home/#{user}/app"
set :branch, "master"
set :start_daemon, '/etc/init.d/go_demo'

task :staging do
  server "stage.adjust.io", :app
end

after 'deploy:update_code', 'go:build'
after "deploy:restart", "go:restart"
after "deploy:start", "go:start"
after "deploy:stop", "go:stop"

namespace :go do
  task :build do
    run "mkdir #{release_path}/bin"
    run "mkdir -p #{release_path}/.go/src/github.com/adeven/gogo"
    run "cd #{release_path} && goem bundle q"
    run "cd #{release_path} &&  goem build #{release_path}/bin/gogo"
  end

  task :start, :roles => :app do
    run "sudo #{start_daemon} start"
  end

  task :restart do
    run "sudo #{start_daemon} restart"
  end

  task :stop do
    run "sudo #{start_daemon} stop"
  end
end

