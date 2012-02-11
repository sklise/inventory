require 'bundler'
Bundler.require

configure do |c|
  # enable :sessions
  set :root, File.dirname(__FILE__)
  set :views, Proc.new{ File.join(root, "app", "views")}
  set :scess, :style => :compact
  set :database, ENV['DATABASE_URL'] || "sqlite3://db/development.sqlite"
  ActiveRecord::Base.include_root_in_json = false

  use Rack::Auth::Basic, "Restricted Area" do |username, password|
    u = ENV['USERNAME'] || 'admin'
    p = ENV['PASSWORD'] || 'admin'
    [username, password] == [u, p]
  end
end

require './app/models'
require './app/routes'
