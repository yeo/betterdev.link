defmodule Betterdev.Web.CollectionController do
  use Betterdev.Web, :controller

  alias Betterdev.Community
  alias Betterdev.Community.Collection

  action_fallback Betterdev.Web.FallbackController

  def index(conn, _params) do
    collections = Community.list_collections(conn.assigns.current_user)
    render(conn, "index.json", collections: collections)
  end

  def create(conn, %{"collection" => collection_params}) do
    with {:ok, %Collection{} = collection} <- Community.create_collection(conn.assigns.current_user, collection_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", collection_path(conn, :show, collection))
      |> render("show.json", collection: collection)
    end
  end

  def show(conn, %{"id" => id}) do
    collection = Community.get_collection!(id)
    render(conn, "show.json", collection: collection)
  end

  def update(conn, %{"id" => id, "collection" => collection_params}) do
    collection = Community.get_collection!(id)

    with {:ok, %Collection{} = collection} <- Community.update_collection(collection, collection_params) do
      render(conn, "show.json", collection: collection)
    end
  end

  def delete(conn, %{"id" => id}) do
    collection = Community.get_collection!(id)
    with {:ok, %Collection{}} <- Community.delete_collection(collection) do
      send_resp(conn, :no_content, "")
    end
  end
end
