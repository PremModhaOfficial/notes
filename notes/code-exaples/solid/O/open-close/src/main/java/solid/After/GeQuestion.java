package solid.After;

public class GeQuestion implements Question {
	public void exeuteGeQuestion() {
		System.out.println("GeQuestion.exeuteGeQuestion()");
	}

	@Override
	public void prosses() {
		exeuteGeQuestion();
	}
}
