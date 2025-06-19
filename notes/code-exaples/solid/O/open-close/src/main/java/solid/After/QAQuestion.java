package solid.After;

public class QAQuestion implements Question {
	public void exeuteFinace() {
		System.out.println("QAQuestion.exeuteFinace()");
	}

	@Override
	public void prosses() {
		exeuteFinace();
	}
}
