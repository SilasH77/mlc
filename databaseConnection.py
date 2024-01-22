import sqlite3

class databaseConnection:

    def __init__(self):

        try:

            # Connect to DB and create a cursor
            sqliteConnection = sqlite3.connect('weeklyPlaylists.db')
            cursor = sqliteConnection.cursor()
            print('DB Init')

            # Close the cursor
            cursor.close()

        # Handle errors
        except sqlite3.Error as error:
            print('Error occurred - ', error)

        # Close DB Connection irrespective of success
        # or failure
        finally:

            if sqliteConnection:
                sqliteConnection.close()
                print('SQLite Connection closed')

#create table of playlists
def createPlaylistTable():
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    # Drop the weeklyplaylist table if already exists.
    cursor.execute("DROP TABLE IF EXISTS weeklyPlaylists")

    # Creating table
    table = """ CREATE TABLE weeklyPlaylists(
                    playlistID VARCHAR(255) NOT NULL,
                    playlistTitle VARCHAR(255) NOT NULL
                ); """

    cursor.execute(table)
    cursor.close()

#add a playlist to the playlist table
def addPlaylistEntry(playlistID: str, playlistTitle: str):
    #connect to table of playlists
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    #insert new playlist
    print("INSERT INTO weeklyPlaylists VALUES ('" + playlistID + "', '" + playlistTitle + "')")
    cursor.execute("INSERT INTO weeklyPlaylists VALUES ('" + playlistID + "', '" + playlistTitle + "')")
    sqliteConnection.commit()
    cursor.close()

#get playlists from playlist table
def getPlaylists():
    # connect to table of playlists
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    #display data from playlist table
    data = cursor.execute('''SELECT * FROM weeklyPlaylists''')
    for row in data:
        print(row)
    cursor.close()

#table containing songs on a particular playlist (need songID and playlistID to identify)
def createSongsTable():
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    # Drop the weeklyplaylist table if already exists.
    cursor.execute("DROP TABLE IF EXISTS trackListings")

    # Creating table
    table = """ CREATE TABLE trackListings(
                        playlistID VARCHAR(255) NOT NULL,
                        trackID VARCHAR(255) NOT NULL,
                        trackTitle VARCHAR(255) NOT NULL
                    ); """

    cursor.execute(table)
    cursor.close()

#add a playlist to the playlist table
def addSongEntry(playlistID: str, trackID: str, trackTitle: str):
    #connect to table of playlists
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    #insert new playlist
    print("INSERT INTO trackListings VALUES ('" + playlistID + "', '" + trackID + "', '" + trackTitle + "')")
    cursor.execute("INSERT INTO trackListings VALUES ('" + playlistID + "', '" + trackID + "', '" + trackTitle + "')")
    sqliteConnection.commit()
    cursor.close()

#get songs from a playlist
def getSongs(playlistID, trackID):
    # connect to table of playlists
    sqliteConnection = sqlite3.connect('spotifyDatabase.db')
    cursor = sqliteConnection.cursor()

    #display data from playlist table
    data = cursor.execute("SELECT trackTitle FROM trackListings WHERE playlistID = '" + playlistID + "' AND trackID = '"  + trackID + "' ")
    for row in data:
        print(row)
    cursor.close()
