package solid.After;

public class AIQuestion implements Question {
	public void exeuteAIQuestion() {
		System.out.println("AIQuestion.exeuteAIQuestion()");
	}

	@Override
	public void prosses() {
		exeuteAIQuestion();
	}

}
