defmodule Betterdev.Repo.Migrations.AddAdminField do
  use Ecto.Migration

  def change do
    alter table(:accounts_users) do
      add :admin, :boolean
    end
  end
end
