package solid.Before;

public class InterviewQuestionprossesor {
	public static void prosses(Question question) {
		if (question instanceof AIQuestion) {
			((AIQuestion) question).exeuteAIQuestion();
		} else if (question instanceof QAQuestion) {
			((QAQuestion) question).exeuteFinace();
		} else if (question instanceof GeQuestion) {
			((GeQuestion) question).exeuteGeQuestion();
		}
	}
}
