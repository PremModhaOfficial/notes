package solid.examples;

import solid.Before.AIQuestion;
import solid.Before.GeQuestion;
import solid.Before.InterviewQuestionprossesor;
import solid.Before.QAQuestion;

public class App {
    public static void main(String[] args) {
        InterviewQuestionprossesor.prosses(new AIQuestion());
        InterviewQuestionprossesor.prosses(new QAQuestion());
        InterviewQuestionprossesor.prosses(new GeQuestion());

        // InterviewQuestionprossesor.prosses(new SomeNewTopic()); will need more effort
        // and less extensible
    }
}
