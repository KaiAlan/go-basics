package main

import (
	"fmt"
)

type Song struct {
	name   string
	artist string
	prev   *Song
	next   *Song
}

type Playlist struct {
	name       string
	creator    string
	head       *Song
	tail       *Song
	nowPlaying *Song
}

func CreatePlaylist(name, creator string) *Playlist {
	return &Playlist{
		name:    name,
		creator: creator,
	}
}

func (p *Playlist) AddSong(name, artist string) error {
	song := &Song{
		name:   name,
		artist: artist,
	}

	if p.head == nil {
		p.head = song
		p.head.next = p.head
		p.head.prev = p.head
	} else {
		p.tail.next = song
		song.prev = p.tail
		song.next = p.head
		p.head.prev = song
	}
	p.tail = song
	return nil
}

func (p *Playlist) ShowAllSong() error {
	if p.head == nil {
		fmt.Println("Playlist is Empty")
		return nil
	}

	currentNode := p.head

	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != p.head {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

func (p *Playlist) StartPlaying() *Song {
	p.nowPlaying = p.head
	return p.nowPlaying
}
func (p *Playlist) NextSong() *Song {
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}
func (p *Playlist) PreviousSong() *Song {
	p.nowPlaying = p.nowPlaying.prev
	return p.nowPlaying
}

func main() {

	playlistName := "myplaylist"
	playlistAuthor := "me"
	myPlaylist := CreatePlaylist(playlistName, playlistAuthor)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Adding songs to the playlist...\n\n")
	myPlaylist.AddSong("Hold on", "Justin Bieber")
	myPlaylist.AddSong("Shape of you", "Ed Sheeran")
	myPlaylist.AddSong("Stubborn Love", "The Lumineers")
	myPlaylist.AddSong("Feels", "Calvin Harris")
	fmt.Println("Showing all songs in playlist...")
	myPlaylist.ShowAllSong()
	fmt.Println()

	myPlaylist.StartPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.NextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.PreviousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.PreviousSong()
	fmt.Println("Changing previous song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

}
