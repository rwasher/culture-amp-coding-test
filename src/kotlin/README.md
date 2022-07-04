# Survey Tool - Kotlin

_Please review the [main README](../../README.md) for overall instructions._

Welcome to the Kotlin starter kit for the survey tool exercise!

Included here is a simple `main.kt` with a `main()` function for invoking the application
as a CLI, and a separate `runAnalysis` function for holding the beginnings of your application
logic. There’s also a lightweight test defined in `MainKtTest.kt` which should
help you drive the beginning of your implementation.

## Prerequisites

- OpenJDK 12+
- Kotlin Compiler 1.3 
- Gradle 6+

## Development
Using an IDE like IntelliJ simplifies development. 
If you want to run things on the command line, you can use the following process. 
### Building the code into a jar

``` sh
kotlinc src/main/kotlin/main.kt -include-runtime -d survey-tool.jar
```

### Testing

``` sh
./gradlew test
```

### Running

``` sh
java -jar survey-tool.jar ../../example-data/survey.csv participation
```
