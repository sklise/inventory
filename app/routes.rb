get '/' do
  erb :main
end

get '/vinyls' do
  content_type :json
  Vinyl.all.to_json(:include => [:author, :label])
end

post '/vinyls/?' do
  content_type :json
  attributes = JSON.parse request.body.read

  # raise attributes.inspect

  @vinyl = Vinyl.new(attributes["vinyl"])
  @vinyl.author = Author.find_or_create_by_name(attributes["author"]["name"])
  @vinyl.label = Label.find_or_create_by_name(attributes["label"]["name"])
  if @vinyl.save
    @vinyl.to_json
  end
end

get '/vinyls/:id' do
  content_type :json
  Vinyl.find(params[:id]).to_json
end

put '/vinyls/:id' do
  content_type :json
  attributes = JSON.parse request.body.read
  
  puts attributes.inspect
  
  @vinyl = Vinyl.find(params[:id])
  @vinyl.update_attributes({
    title: attributes["title"],
    year: attributes["year"],
    records: attributes["records"]
  })

  if @vinyl.author.nil?
    @vinyl.author = Author.find_or_create_by_name(attributes["author"]["name"])
  else
    @vinyl.author.update_attributes name: attributes["author"]["name"]
  end

  if @vinyl.label.nil?
    @vinyl.label = Label.find_or_create_by_name(attributes["label"]["name"])
  else
    @vinyl.label.update_attributes name: attributes["label"]["name"]
  end

  @vinyl.save
  @vinyl.to_json
end

delete '/vinyls/:id' do |id|
  content_type :json
  @vinyl = Vinyl.find(params[:id])
  @vinyl.destroy
  {:head => :ok}.to_json
end