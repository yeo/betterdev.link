defmodule Betterdev.Web.CollectionView do
  use Betterdev.Web, :view
  alias Betterdev.Web.CollectionView

  def render("index.json", %{collections: collections}) do
    %{data: render_many(collections, CollectionView, "collection.json")}
  end

  def render("show.json", %{collection: collection}) do
    %{data: render_one(collection, CollectionView, "collection.json")}
  end

  def render("collection.json", %{collection: collection}) do
    %{id: collection.id,
      name: collection.name,
      user_id: collection.user_id}
  end
end
