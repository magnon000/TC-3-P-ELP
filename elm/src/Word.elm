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

type alias Model =
  {
    modelStatus : Status,
    score : Float,
    userInput : String
  }

type Status
  = FailureWords String
  | FailureRand
  | FailureRecuperationDuMotChoisi
  | FailureAPI String
  | OtherFailure
  | Loading
  | AllWords String
  | OneWord String
  | WordWithData String WordData
  | GuessingPhase String WordData
  | GivenUp String WordData
  --| HintGiven String WordData

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
  ( modelFromStatus Loading
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
  | NewUserGuess String
  | UserGivingUp
  --| UserAskingHint


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          ((modelFromStatus (AllWords fullText)), Random.generate GotRand (Random.int 1 999))

        Err error ->
          wordListErrorHandler error

    GotRand id ->
        case model.modelStatus of
            AllWords fullText ->
                ((modelFromStatus (OneWord (getWordAtIndex id (words fullText)))), getWordJson (getWordAtIndex id (words fullText)))
            _ ->
                ((modelFromStatus (FailureRand)), Cmd.none)
    GotEverything result ->
      case model.modelStatus of
            OneWord theWord ->
                case result of
                  Ok data ->
                    ((modelFromStatus (WordWithData theWord data)), Cmd.none)
                  
                  Err error ->
                    jsonErrorHandler error

            _ ->
              ((modelFromStatus (FailureRecuperationDuMotChoisi)), Cmd.none)
    
    NewUserGuess guess ->
      case model.modelStatus of
        WordWithData theWord data ->
          ({ model | modelStatus = (GuessingPhase theWord data), userInput = guess }, Cmd.none)
        
        GuessingPhase theWord data ->
          ({ model | userInput = guess }, Cmd.none)

        _ ->
          ((modelFromStatus (OtherFailure), Cmd.none))
    
    UserGivingUp ->
      case model.modelStatus of
        WordWithData theWord data ->
          giveUpModel model theWord data
        
        GuessingPhase theWord data ->
          giveUpModel model theWord data

        _ ->
          ((modelFromStatus (OtherFailure), Cmd.none))

    




-- VIEW


view : Model -> Html Msg
view model =
  case model.modelStatus of
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
      viewNormal model word data
    
    GuessingPhase word data ->
      viewNormal model word data
    
    GivenUp word data->
      viewNormal model word data





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




viewNormal : Model -> String -> WordData -> Html Msg
viewNormal model word data = 
  div [] [
    div [style "text-align" "center"] [
      h5 [] [text "MONTEAGUDO Diego & MA Longrui's Word Guesser Game"]
      ,viewValidation model
      ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
        input [ placeholder "Your guess here", value model.userInput, onInput NewUserGuess ] []
      ]
      --, button [ onClick UserAskingHint ] [ text "HINT (-0.5 points)" ]
      ,hr [] []
    ],
    div [] [
    wordDataToHtml data
    ]
  ]


viewValidation : Model -> Html Msg
viewValidation model = case model.modelStatus of
  GuessingPhase word data ->
    if word == model.userInput then
      div [] [
        div[] [ h1 [ style "padding-left" "30px", style "color" "green"] [ text ":)"]]
        ,hr [] []
        ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
          h3 [ style "color" "green"] [ text (String.append "YES !!! The word was indeed \"" (String.append word "\" !"))]
        ]
      ]
    else
      div [] [
        div[] [ h1 [ style "padding-left" "30px"] [ text "Try to guess the word"]
        ,hr [] []
        , button [ onClick UserGivingUp ] [ text "I give up" ]
        ]
      ]

  WordWithData word data ->
    div[] [ h1 [ style "padding-left" "30px"] [ text "Try to guess the word"]
    ,hr [] []
    , button [ onClick UserGivingUp ] [ text "I give up" ]
    ]
  
  GivenUp word data ->
    div [] [
      div[] [ h1 [ style "padding-left" "30px"] [ text ":|"]]
      ,hr [] []
      ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
        h3 [ style "color" "red"] [ text (String.append "The word was \"" (String.append word "\"..."))]
      ]
    ]
  
  _ ->
    div [] []








modelFromStatus : Status -> Model
modelFromStatus status = (Model status 0 "")


giveUpModel : Model -> String -> WordData -> (Model, Cmd Msg)
giveUpModel model word wordData =
  if model.score == 0 then
    ({ model | modelStatus = (GivenUp word wordData) }, Cmd.none)
  else
    ({ model | score = (model.score - 1), modelStatus = (GivenUp word wordData) }, Cmd.none)








wordListErrorHandler : Http.Error -> (Model, Cmd Msg)
wordListErrorHandler error = case error of
  Http.BadUrl _ ->
    ((modelFromStatus (FailureWords "Bad Url")), Cmd.none)
  Http.Timeout ->
    ((modelFromStatus (FailureWords "Timeout")), Cmd.none)
  Http.NetworkError ->
    ((modelFromStatus (FailureWords "Network Error")), Cmd.none)
  Http.BadStatus _ ->
    ((modelFromStatus (FailureWords "Bad Status")), Cmd.none)
  Http.BadBody _ ->
    ((modelFromStatus (FailureWords "Bad Body")), Cmd.none)

jsonErrorHandler : Http.Error -> (Model, Cmd Msg)
jsonErrorHandler error = case error of
  Http.BadUrl _ ->
      ((modelFromStatus (FailureAPI "Bad Url")), Cmd.none)
  Http.Timeout ->
      ((modelFromStatus (FailureAPI "Timeout")), Cmd.none)
  Http.NetworkError ->
      ((modelFromStatus (FailureAPI "Network Error")), Cmd.none)
  Http.BadStatus _ ->
      ((modelFromStatus (FailureAPI "Bad Status")), Cmd.none)
  Http.BadBody problem ->
      ((modelFromStatus (FailureAPI (String.append "Bad Body" problem))), Cmd.none)




-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



