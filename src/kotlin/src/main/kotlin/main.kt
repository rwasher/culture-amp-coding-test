import java.io.OutputStream
import java.io.PrintStream

fun main(args: Array<String>) {
    if (args.size != 2) {
        println("Usage: java -jar survey-tool.jar <file path> <analysis mode>")
        throw IllegalArgumentException()
    }
    val file = args[0]
    val mode = args[1]

    println("Run '$mode' analysis on '$file'\n")

    runAnalysis(mode, file, System.out, System.err)
}

fun runAnalysis(mode: String, file: String, outputStream: OutputStream, errorStream: OutputStream) {
    // Your implementation goes here
    val result = ""
    PrintStream(outputStream).print(result)
}
