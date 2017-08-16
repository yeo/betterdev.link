defmodule Betterdev.Repo.Migrations.AddBotChannelToUser do
  use Ecto.Migration

  def change do
    alter table(:accounts_users) do
      add :channel, :string
      add :channel_user_id, :string

      create index(:community_link_tags, [:link_id, :tag_id], unique: true)
    end
  end
end
