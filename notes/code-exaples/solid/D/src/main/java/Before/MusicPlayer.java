package Before;

public class MusicPlayer {
	private NetworkMusicLoader networkMusicLoader = new NetworkMusicLoader();

	public MusicPlayer() {

	}

	public void play() {
		networkMusicLoader.load();
	}

}
