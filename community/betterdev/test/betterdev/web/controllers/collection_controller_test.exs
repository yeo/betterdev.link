defmodule Betterdev.Web.CollectionControllerTest do
  use Betterdev.Web.ConnCase

  alias Betterdev.Community
  alias Betterdev.Community.Collection

  @create_attrs %{name: "some name", user_id: 42}
  @update_attrs %{name: "some updated name", user_id: 43}
  @invalid_attrs %{name: nil, user_id: nil}

  def fixture(:collection) do
    {:ok, collection} = Community.create_collection(@create_attrs)
    collection
  end

  setup %{conn: conn} do
    {:ok, conn: put_req_header(conn, "accept", "application/json")}
  end

  test "lists all entries on index", %{conn: conn} do
    conn = get conn, collection_path(conn, :index)
    assert json_response(conn, 200)["data"] == []
  end

  test "creates collection and renders collection when data is valid", %{conn: conn} do
    conn = post conn, collection_path(conn, :create), collection: @create_attrs
    assert %{"id" => id} = json_response(conn, 201)["data"]

    conn = get conn, collection_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "name" => "some name",
      "user_id" => 42}
  end

  test "does not create collection and renders errors when data is invalid", %{conn: conn} do
    conn = post conn, collection_path(conn, :create), collection: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "updates chosen collection and renders collection when data is valid", %{conn: conn} do
    %Collection{id: id} = collection = fixture(:collection)
    conn = put conn, collection_path(conn, :update, collection), collection: @update_attrs
    assert %{"id" => ^id} = json_response(conn, 200)["data"]

    conn = get conn, collection_path(conn, :show, id)
    assert json_response(conn, 200)["data"] == %{
      "id" => id,
      "name" => "some updated name",
      "user_id" => 43}
  end

  test "does not update chosen collection and renders errors when data is invalid", %{conn: conn} do
    collection = fixture(:collection)
    conn = put conn, collection_path(conn, :update, collection), collection: @invalid_attrs
    assert json_response(conn, 422)["errors"] != %{}
  end

  test "deletes chosen collection", %{conn: conn} do
    collection = fixture(:collection)
    conn = delete conn, collection_path(conn, :delete, collection)
    assert response(conn, 204)
    assert_error_sent 404, fn ->
      get conn, collection_path(conn, :show, collection)
    end
  end
end
