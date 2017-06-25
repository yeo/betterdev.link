defmodule Betterdev.Web.LinkController do
  use Betterdev.Web, :controller

  alias Betterdev.Community
  alias Betterdev.Accounts
  alias Betterdev.Community.Link

  action_fallback Betterdev.Web.FallbackController

  def index(conn, params) do
    {postings, kerosene} = Community.list_links(params)
    render(conn, "index.json", links: postings, kerosene: kerosene)
  end

  def create(conn, %{"link" => link_params}) do
    #with {:ok, %Link{} = link} <- Community.create_link(link_params, conn.assigns.current_user) do
    with {:ok, %Link{} = link} <- Community.user_post_link(conn.assigns.current_user, link_params) do
      conn
      |> put_status(:created)
      |> put_resp_header("location", link_path(conn, :show, link))
      |> render("show.json", link: link)
    end
  end

  def show(conn, %{"id" => id}) do
    link = Community.get_link!(id)
    render(conn, "show.json", link: link)
  end

  def update(conn, %{"id" => id, "link" => link_params}) do
    link = Community.get_link!(id)

    with {:ok, %Link{} = link} <- Community.update_link(link, link_params) do
      render(conn, "show.json", link: link)
    end
  end

  def delete(conn, %{"id" => id}) do
    link = Community.get_link!(id)
    with {:ok, %Link{}} <- Community.delete_link(link) do
      send_resp(conn, :no_content, "")
    end
  end
end
