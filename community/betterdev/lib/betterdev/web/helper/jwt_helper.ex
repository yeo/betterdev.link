defmodule Betterdev.JWTHelpers do
  import Joken, except: [verify: 1]

  @doc """
  use for future verification, eg. on socket connect
  """
  def verify(jwt) do
    verify
    |> with_json_module(Poison)
    |> with_compact_token(jwt)
    |> Joken.verify
  end

  @doc """
  use for verification via plug
  issuer should be our auth0 domain
  app_metadata must be present in id_token
  """
  def verify do
    %Joken.Token{}
    |> with_json_module(Poison)
    |> with_signer(hs256(config[:app_secret]))
    |> with_validation("aud", &(&1 == config[:app_id]))
    |> with_validation("exp", &(&1 > current_time))
    |> with_validation("iat", &(&1 <= current_time))
    |> with_validation("iss", &(&1 == config[:app_baseurl]))
  end

  @doc """
  Create token from client id and secret
  Used for unit tests
  """
  def create_bearer_token(auth_scopes, config_items \\ %{:signer => :app_secret, :aud => :app_id}) do
    %Joken.Token{claims: auth_scopes}
    |> with_json_module(Poison)
    |> with_signer(hs256(config[config_items[:signer]]))
    |> with_aud(config[config_items[:aud]])
    |> with_iat
    |> with_iss(config[:app_baseurl])
    |> with_exp(current_time + 86_400)
    |> sign
    |> get_compact
  end

  @doc """
  Return error message for `on_error`
  """
  def error(conn, _msg) do
    {conn, %{:errors => %{:detail => "unauthorized"}}}
  end

  defp config, do: Application.get_env(:betterdev, :auth0)
end
