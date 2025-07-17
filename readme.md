## Support for i.MX RT microcontrollers.

Embedded Go supports the i.MX RT106x family. You may know it as Teensy 4 which is a family of the popular development boards based on the i.MX RT1062.

### Getting started

1. Install the Embedded Go toolchain.

   ```sh
   go install github.com/embeddedgo/dl/go1.24.5-embedded@latest
   go1.24.5-embedded download
   ```

2. Install egtool.

   ```sh
   go install github.com/embeddedgo/tools/egtool@latest
   ```

3. Create a project directory containing the `main.go` file with your first Go program for Teensy 4.x.

   ```go
   package main

   import (
   	"time"

   	"github.com/embeddedgo/imxrt/devboard/teensy4/board/leds"
   )

   func main() {
   	for {
   		leds.User.Toggle()
   		time.Sleep(time.Second/2)
   	}
   }
   ```

4. Initialize your project

   ```sh
   go mod init firstprog
   go mod tidy
   ```

5. Copy the `go.env` file suitable for your board (here is one for [Teensy](https://github.com/embeddedgo/imxrt/tree/master/devboard/teensy4/examples/go.env) and another one for a [board with 4 MB flash](https://github.com/embeddedgo/imxrt/tree/master/devboard/fet1061/examples/go.env)).

6. Compile your first program

   ```sh
   export GOENV=go.env
   go build
   ```

   or

   ```sh
   GOENV=go.env go build
   ```

7. Connect your Teensy to your computer and press the onboard button.

8. Load and run.

   ```sh
   egtool load
   ```

### Examples

See more example code for [supported develompent boards](devboard).