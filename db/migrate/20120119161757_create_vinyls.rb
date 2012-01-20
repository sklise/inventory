class CreateVinyls < ActiveRecord::Migration
  def change
    create_table :vinyls do |t|
      t.string :title
      t.integer :author_id
      t.integer :year
      t.integer :label_id
      t.integer :size
      t.integer :records
    end
  end
end
