class Vinyl < ActiveRecord::Base
  belongs_to :author
  belongs_to :label
end

class Author < ActiveRecord::Base
  has_many :vinyls
  
  def to_s
    name
  end
end

class Label < ActiveRecord::Base
  has_many :vinyls
end