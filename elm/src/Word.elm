module Word exposing (..)

import Browser
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field)
import String exposing (words)
import Random
import Debug exposing (..)


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
  = FailureWords String
  | FailureRand
  | FailureRecuperationDuMotChoisi
  | FailureAPI String
  | OtherFailure
  | Loading
  | AllWords String
  | OneWord String
  | WordWithData String WordData

type alias WordData =
  { senses : List Sens
  }

type alias Sens =
  { usages : List Usage
  }

type alias Usage =
  { wordType : String
  , definitions : List String
  }


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
  | GotEverything (Result Http.Error WordData)


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          (AllWords fullText, Random.generate GotRand (Random.int 1 999))

        Err error ->
          case error of
            Http.BadUrl _ ->
              (FailureWords "Bad Url", Cmd.none)
            Http.Timeout ->
              (FailureWords "Timeout", Cmd.none)
            Http.NetworkError ->
              (FailureWords "Network Error", Cmd.none)
            Http.BadStatus _ ->
              (FailureWords "Bad Status", Cmd.none)
            Http.BadBody _ ->
              (FailureWords "Bad Body", Cmd.none)
    GotRand id ->
        case model of
            AllWords fullText ->
                (OneWord (getWordAtIndex id (words fullText)), getWordJson (getWordAtIndex id (words fullText)))
            _ ->
                (FailureRand, Cmd.none)
    GotEverything result ->
      case model of
            OneWord theWord ->
                case result of
                  Ok data ->
                    (WordWithData theWord data, Cmd.none)
                  
                  Err error ->
                    case error of
                      Http.BadUrl _ ->
                        (FailureAPI "Bad Url", Cmd.none)
                      Http.Timeout ->
                        (FailureAPI "Timeout", Cmd.none)
                      Http.NetworkError ->
                        (FailureAPI "Network Error", Cmd.none)
                      Http.BadStatus _ ->
                        (FailureAPI "Bad Status", Cmd.none)
                      Http.BadBody problem ->
                        (FailureAPI (String.append "Bad Body" problem), Cmd.none)

            _ ->
              (FailureRecuperationDuMotChoisi, Cmd.none)
    
    -- UserGuess guess ->
    --   case model of ->
    --     WordWithData theWord data ->
    --       (GuessingPhase guess theWord data, Cmd.none)
        
    --     _ ->
    --       (OtherFailure, Cmd.none)

    

            



-- VIEW


view : Model -> Html Msg
view model =
  case model of
    FailureWords reason ->
      text (String.append "Erreur HTTP récupération des mots : " reason)
    
    FailureRand ->
      text "Échec de la randomisation."
    
    FailureRecuperationDuMotChoisi ->
      text "Echec récupération du mot qui avait été choisi aléatoirement"
    
    FailureAPI reason ->
      text (String.append "Erreur HTTP récupération du json : " reason)
    
    OtherFailure ->
      text "General unspecific failure"

    Loading ->
      text "Loading..."

    AllWords fullText ->
      pre [] [ text fullText ]
    
    OneWord word ->
      h1 
      [ style "top" "30px" --marche pas
      , style "left" "50px"] --marche pas
      [ text word ]
    
    WordWithData word data ->
      div [] [
        h1 [] [text "Try to guess the word"]
        ,hr [] []
        ,hr [] []
        ,wordDataToHtml data
      ]
    
    --GuessingPhase guess word wordData ->




getWordAtIndex : Int -> List String -> String
getWordAtIndex index liste = case liste of
    [] -> "ERROR (not the word error)"
    (x::xs) -> case index of
        0 -> x
        _ -> getWordAtIndex (index-1) xs --dit erreur sur vscode mais fonctionne...




getWordJson : String -> Cmd Msg
getWordJson word =
  Http.get
    { url = (String.append "https://api.dictionaryapi.dev/api/v2/entries/en/" word)
    , expect = Http.expectJson GotEverything jsonDecoder
    }


jsonDecoder : Decoder WordData
jsonDecoder =
  Json.Decode.map WordData (Json.Decode.list senseDecoder)

senseDecoder : Decoder Sens
senseDecoder =
  Json.Decode.map Sens
    (field "meanings" (Json.Decode.list usageDecoder))

usageDecoder : Decoder Usage
usageDecoder =
  Json.Decode.map2 Usage
    (field "partOfSpeech" Json.Decode.string)
    (field "definitions" (Json.Decode.list (field "definition" Json.Decode.string)))






getFirstSense : List Sens -> Sens
getFirstSense senses = case senses of
    [] -> (Sens [])
    (x::xs) -> x

getFirstUsage : List Usage -> Usage
getFirstUsage usages = case usages of
    [] -> (Usage "" [])
    (x::xs) -> x







---TEST 


-- wordDataToHtml : WordData -> Html msg
-- wordDataToHtml wordData =
--   ul [] (List.map sensToHtml wordData.senses)

-- sensToHtml : Sens -> Html msg
-- sensToHtml sens =
--   li [] [ (text "meaning :") :: List.map usageToHtml sens.usages ]

-- usageToHtml : Usage -> Html msg
-- usageToHtml usage =
--   ul []
--     [ li [] [ text (String.append "word type : " usage.wordType) ]
--     , ol [] [(List.map definitionToHtml usage.definitions)]
--     ]

-- definitionToHtml : String -> Html msg
-- definitionToHtml definition =
--   li [] [ text definition ]

wordDataToHtml : WordData -> Html msg
wordDataToHtml wordData =
  ul [] (List.map sensToHtml wordData.senses)

sensToHtml : Sens -> Html msg
sensToHtml sens =
  li [style "padding-top" "10px"] ((h3 [] [text "Possible meaning :"]) :: List.map usageToHtml sens.usages)

usageToHtml : Usage -> Html msg
usageToHtml usage =
  ul [] [
    li [] [ h4 [] [(text (String.append "Word type : " usage.wordType))]],
    ol [] (List.map definitionToHtml usage.definitions)
  ]

definitionToHtml : String -> Html msg
definitionToHtml definition =
  li [style "padding-top" "2px"] [ text definition ]











-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



