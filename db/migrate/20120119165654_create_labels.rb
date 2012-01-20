class CreateLabels < ActiveRecord::Migration
  def change
    create_table :labels do |t|
      t.string :name
    end
  end
end
