package After;

public class MusicPlayer {
	private MusicLoader loader;

	public MusicPlayer(MusicLoader loader) {
		this.loader = loader;
	}

	public void play() {
		loader.load();
	}

}
