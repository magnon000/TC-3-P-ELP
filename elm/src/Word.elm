module Word exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, map4, field, int, string)
import String exposing (words)
import Random


import Browser
import Html exposing (Html, text, pre)
import Http



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type Model
  = Failure
  | Loading
  | AllWords String
  | OneWord String


init : () -> (Model, Cmd Msg)
init _ =
  ( Loading
  , Http.get
      { url = "../static/words.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | GotRand Int


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          (AllWords fullText, Random.generate GotRand (Random.int 1 999))

        Err _ ->
          (Failure, Cmd.none)
    GotRand id ->
        case model of
            AllWords fullText ->
                (OneWord (getWordAtIndex id (words fullText)), Cmd.none)
            _ ->
                (Failure, Cmd.none)

            



-- VIEW


view : Model -> Html Msg
view model =
  case model of
    Failure ->
      text "Liste de mots introuvable."

    Loading ->
      text "Loading..."

    AllWords fullText ->
      pre [] [ text fullText ]
    
    OneWord word ->
      h1 
      [ style "top" "30px" --marche pas
      , style "left" "50px"] --marche pas
      [ text word ]



getWordAtIndex : Int -> List String -> String
getWordAtIndex index liste = case liste of
    [] -> "ERROR (not the word error)"
    (x::xs) -> case index of
        0 -> x
        _ -> getWordAtIndex (index-1) xs --dit erreur sur vscode mais fonctionne...


























-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none


