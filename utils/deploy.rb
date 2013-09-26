set :application, "datatable"
set :repository, "."
set :scm, :none
set :user, "server_username"
set :deploy_to, "/home/#{user}/path/to/#{application}"
set :deploy_via, :copy
set :current_path, "/path/to/current/#{application}"
set :copy_exclude, ["*.go",".gitignore",".git","utils","Makefile","Capfile","config"]
set :build_script, "go build -o #{application}"
set :copy_strategy, :export

server "server", :app, :web, :db, :primary => true

namespace :remote do
  task :create_releases do
    run "mkdir -p #{deploy_to}/releases"
  end
end

namespace :deploy do
  task :migrate do
    #no-op
  end
  task :finalize_update do
    #no-op
  end
  task :start do
    run "start_script" #/etc/init.d/datatable start
  end
  task :stop do
    run "stop_script" #/etc/init.d/datatable stop
  end
  task :restart do
    run "restart_command" #/etc/init.d/datatable restart
  end
end

after "deploy:setup", "remote:create_releases"
