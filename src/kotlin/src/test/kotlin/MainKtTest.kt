import org.junit.After
import org.junit.Assert.*
import org.junit.Before
import org.junit.Test
import java.io.ByteArrayOutputStream
import java.io.OutputStream
import java.io.PrintStream


class MainKtTest {
    @Test
    fun testParticipation() {
        val output = ByteArrayOutputStream()
        val outputStream = PrintStream(output)
        val error = ByteArrayOutputStream()
        val errorStream = PrintStream(error)
        val file = "example-data/survey.csv"
        val mode = "participation"

        runAnalysis(mode, file, outputStream, errorStream)

        val expectedParticipationResult =
            """
				|Participation
				|Participants: 6
				|Submitted: 5 (83.3%)
            """.trimMargin()
        assertEquals(expectedParticipationResult, output.toString());
    }
}