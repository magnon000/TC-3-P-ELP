module Word exposing (..)

import Browser
import Html exposing (..)
import Html exposing (Html, text, pre)
import Html.Attributes exposing (..)
import Html.Events exposing (..)
import Http
import Json.Decode exposing (Decoder, field)
import String exposing (words)
import Random



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
    score : Int,
    userInput : String
  }

-- records utilisés

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
  | Won String WordData
  --| HintGiven String WordData


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
  | CorrectlyGuessed
  | PlayAgain
  --| UserAskingHint


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    -- Texte obtenu : Préparer un nombre aléatoire
    GotText result ->
      case result of
        Ok fullText ->
          ({ model | modelStatus = (AllWords fullText)}, (Random.generate GotRand (Random.int 1 999)))

        Err error ->
          wordListErrorHandler error
    
    -- Nombre aléatoire obtenu : choisi le mot correspondant et chercher les définitions
    GotRand id ->
        case model.modelStatus of
            AllWords fullText ->
                ({ model | modelStatus = OneWord (getWordAtIndex id (words fullText))}, getWordJson (getWordAtIndex id (words fullText)))
            _ ->
                ((modelFromStatus (FailureRand)), Cmd.none)
    
    -- Tout est prêt : rendre possible l'affichage des définitions
    GotEverything result ->
      case model.modelStatus of
            OneWord theWord ->
                case result of
                  Ok data ->
                    ({ model | modelStatus = (WordWithData theWord data)}, Cmd.none)
                  
                  Err error ->
                    jsonErrorHandler error

            _ ->
              ((modelFromStatus (FailureRecuperationDuMotChoisi)), Cmd.none)
    
    NewUserGuess guess ->
      case model.modelStatus of
        WordWithData theWord data ->
          ({ model | modelStatus = (GuessingPhase theWord data), userInput = guess }, Cmd.none)
        
        GuessingPhase theWord data ->
          correctTest model guess theWord data

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
    
    CorrectlyGuessed ->
      case model.modelStatus of
        GuessingPhase theWord data ->
          ({ model | score = (model.score + 1), modelStatus = (Won theWord data)}, Cmd.none)
        
        _ ->
          ((modelFromStatus (OtherFailure), Cmd.none))
    
    PlayAgain ->
      newGameModel model

    




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
      bait model

    AllWords fullText ->
      bait model
    
    OneWord word ->
      bait model
    
    -- phase initiale : le joueur peut commencer à écrire
    WordWithData word data ->
      div [] [
        div [style "text-align" "center"] [
            partieHaute "Try to guess the word" model
            ,button [ onClick UserGivingUp ] [ text "Give up ( -1 point)" ]
          ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
            input [ placeholder "Your guess here", value model.userInput, onInput NewUserGuess ] []
          ,hr [] []
          ]
        ],
        div [] [
          wordDataToHtml data
        ]
      ]
    
    -- le joueur a déjà écrit au moins 1 fois
    GuessingPhase word data ->
      div [] [
        div [style "text-align" "center"] [
          partieHaute "Try to guess the word" model
          ,button [ onClick UserGivingUp ] [ text "Give up ( -1 point)" ]
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
    
    -- abandon
    GivenUp word data->
      div [] [
        div [style "text-align" "center"] [
          partieHaute ":|" model
          ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
            h3 [ style "color" "red"] [ text (String.append "The word was \"" (String.append word "\"..."))]
          ]
          --, button [ onClick UserAskingHint ] [ text "HINT (-0.5 points)" ]
          ,button [ onClick PlayAgain ] [ text "Continue to play" ]
          ,hr [] []
        ],
        div [] [
          wordDataToHtml data
        ]
      ]

    -- victoire
    Won word data ->
      div [] [
        div [style "text-align" "center"] [
            partieHaute ":)" model
            ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
              h3 [ style "color" "green"] [ text (String.append "YES !!! The word was indeed \"" (String.append word "\" !"))]
            ]
        --, button [ onClick UserAskingHint ] [ text "HINT (-0.5 points)" ]
        ,button [ onClick PlayAgain ] [ text "Continue to play" ]
        ,hr [] []
        ]
        ,div [] [
          wordDataToHtml data
        ]
      ]






-- a mettre dans module séparé

-- DECODEURS

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



-- FONCTIONS D'AFFICHAGE DU JSON

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



getWordJson : String -> Cmd Msg
getWordJson word =
  Http.get
    { url = (String.append "https://api.dictionaryapi.dev/api/v2/entries/en/" word)
    , expect = Http.expectJson GotEverything jsonDecoder
    }

getWordAtIndex : Int -> List String -> String
getWordAtIndex index liste = case liste of
    [] -> "ERROR (not the word error)"
    (x::xs) -> case index of
        0 -> x
        _ -> getWordAtIndex (index-1) xs --dit erreur sur vscode mais fonctionne...









-- AIDE A L'AFFICHAGE

-- Elements toujours présents en haut de la page
partieHaute : String -> Model -> Html Msg
partieHaute contenu model =
    div [] [
      h5 [] [text "MONTEAGUDO Diego & MA Longrui's Word Guesser Game"]
      ,div [] [
        div[] [ h1 [ style "padding-left" "30px"] [ text contenu]]
      ]
      ,h3 [] [ text (String.append "Score: " (String.fromInt model.score))]
      ,hr [] []
    ]


-- Elements qui ne font rien pendant les phases de chargement (fluidité visuelle)
bait : Model -> Html Msg
bait model =
  div [] [
    div [style "text-align" "center"] [
      partieHaute "Loading..." model
      ,button [ ] [ text "Give up ( -1 point)" ]
      ,div [style "padding-bottom" "10px", style "padding-top" "10px"] [
        input [ placeholder "Your guess here"] []
        ,hr [] []
      ]
    ]
  ]





-- FONCTIONS MODELES

-- Création rapide d'un modèle avec un status précis
modelFromStatus : Status -> Model
modelFromStatus status = (Model status 0 "")

-- Recommencer depuis le modèle initial en gardant le score
newGameModel : Model -> (Model, Cmd Msg)
newGameModel model = ((Model Loading model.score ""), Http.get { 
  url = "../static/words.txt" , expect = Http.expectString GotText })

-- Pour ne pas aller en dessous de 0 en score
giveUpModel : Model -> String -> WordData -> (Model, Cmd Msg)
giveUpModel model word wordData =
  if model.score == 0 then
    ({ model | modelStatus = (GivenUp word wordData) }, Cmd.none)
  else
    ({ model | score = (model.score - 1), modelStatus = (GivenUp word wordData) }, Cmd.none)


-- Gestion de la réussite
correctTest : Model -> String -> String -> WordData -> (Model, Cmd Msg)
correctTest model guess word wordData =
  if word == guess then
    ({ model | userInput = guess, score = model.score + 1, modelStatus = (Won word wordData) }, Cmd.none)
  else
    ({model | userInput = guess}, Cmd.none)






-- GESTION DES ERREURS

-- Erreurs de chargement des mots
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

--Erreurs de chargement du json
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



