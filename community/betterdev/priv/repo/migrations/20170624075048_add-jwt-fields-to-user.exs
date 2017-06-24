defmodule :"Elixir.Betterdev.Repo.Migrations.AddJwtFieldsToUser" do
  use Ecto.Migration

  def change do
    alter table(:accounts_users) do
      add :jwt_sub, :integer
      add :jwt_aud, :string
    end
  end
end
