defmodule Betterdev.Tasks.Sync do
  alias  Betterdev.Community
  alias Betterdev.Community.Collection
  alias Betterdev.Community.Link
  alias Betterdev.Repo

  def retag do
    links = Repo.all(Link)
    Enum.map(links, &(one(&1)))
  end

  defp one(link) do
    IO.inspect link
    link = link |> Repo.preload(:tags) |> Repo.preload(:collections)
    Community.post_process_link(&1)
  end
end
