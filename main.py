import spotipy
from dotenv import load_dotenv
from spotipy.oauth2 import SpotifyOAuth
import os

#pip install python-dotenv, automatically load environment files
load_dotenv()
scope = "playlist-modify-public"
username = "qn6gi7wxga2uu48vv8lzkt8jg"

client_id = os.getenv("SPOTIPY_CLIENT_ID")
client_secret = os.getenv("SPOTIPY_CLIENT_SECRET")

token = SpotifyOAuth(scope=scope,username=username)
spotifyObject = spotipy.Spotify(auth_manager= token)

#this function will create a playlist on spotify using a specified list of songs
def create_playlist(token):
    playlist_name = input("Enter a playlist name: ")
    playlist_description = input("Enter a playlist description: ")

    spotifyObject.user_playlist_create(user=username,name=playlist_name,public=True,
                                      description=playlist_description)


    user_input = input("What song would you like on this playlist: ")

    #find a specified song, and select the top result
    song_list=[]
    result = spotifyObject.search(q=user_input)
    json_result = result["tracks"]["items"][0]["uri"]
    song_list.append(json_result)

    #select playlist from profile
    playlistResults = spotifyObject.user_playlists(user=username)
    playlist = playlistResults["items"][0]["id"]

    #add songs to playlist
    spotifyObject.user_playlist_add_tracks(user=username, playlist_id=playlist,tracks = song_list)

#call create playlist
create_playlist(token)


'''all below is example runs of different functionalities for the database, 
the actual values given will be determined by user input from the UI

#instantiate a connection
sql = databaseConnection
sql.databaseConnection()

#create a table to store references to individual playlists
sql.createPlaylistTable()

#add an entry to the table of playlists
sql.addPlaylistEntry('12', 'Songs to Cry Naked in the Shower to.')
#display playlists
sql.getPlaylists()

#create a table to contain songs to associate with playlists
sql.createSongsTable()

#add an entry to songs
sql.addSongEntry('12', '3', 'Mr. Blue Sky')
#display songs in a playlist
sql.getSongs('12','3')
'''
