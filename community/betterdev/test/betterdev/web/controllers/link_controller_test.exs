defmodule Betterdev.Web.LinkControllerTest do
  use Betterdev.Web.ConnCase

  alias Betterdev.Community
  alias Betterdev.Community.Link

  @create_attrs %{title: "some title", uri: "some uri"}
  @update_attrs %{title: "some updated title", uri: "some updated uri"}
  @invalid_attrs %{title: nil, uri: nil}

  def fixture(:link) do
    {:ok, link} = Community.create_link(@create_attrs)
    link
  end

  setup %{conn: conn} do
    {:ok, conn: put_req_header(conn, "accept", "application/json")}
  end

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, link_path(conn, :index)
    assert json_response(conn, 200)["data"] == []
  end

  test "creates link and renders link when data is valid", %{conn: conn} do
    conn = post conn, link_path(conn, :create), link: @create_attrs
    assert %{"id" => id} = json_response(conn, 201)["data"]

    conn = get conn, link_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "title" => "some title",
      "uri" => "some uri"}
  end

  test "does not create link and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, link_path(conn, :create), link: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "updates chosen link and renders link when data is valid", %{conn: conn} do
    %Link{id: id} = link = fixture(:link)
    conn = put conn, link_path(conn, :update, link), link: @update_attrs
    assert %{"id" => ^id} = json_response(conn, 200)["data"]

    conn = get conn, link_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "title" => "some updated title",
      "uri" => "some updated uri"}
  end

  test "does not update chosen link and renders errors when data is invalid", %{conn: conn} do
    link = fixture(:link)
    conn = put conn, link_path(conn, :update, link), link: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "deletes chosen link", %{conn: conn} do
    link = fixture(:link)
    conn = delete conn, link_path(conn, :delete, link)
    assert response(conn, 204)
    assert_error_sent 404, fn ->
      get conn, link_path(conn, :show, link)
    end
  end
end
