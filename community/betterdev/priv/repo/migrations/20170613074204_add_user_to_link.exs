defmodule Betterdev.Repo.Migrations.AddUserToLink do
  use Ecto.Migration

  def change do
    alter table(:community_links) do
      add :user_id, :integer
      add :picture, :string
      add :status, :string
      add :description, :string
    end
  end
end
