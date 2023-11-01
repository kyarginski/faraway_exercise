#!/usr/bin/expect

# Client app starts
spawn -noecho go run ./cmd/client -addr=127.0.0.1:8088 connect

expect {
    "Received PoW task:" {
        puts "\nReceived PoW task: {\n"
        puts "$expect_out(0,string)"
        puts "}"
    }
}

expect "Steps:"
puts "$expect_out(0,string)"

expect {
    "Found PoW solution:" {
        puts "\nFound PoW solution: {\n"
        puts "$expect_out(0,string)"
        puts "}"
    }
}

expect {
    "Server response:" {
        puts "\nServer response: {\n"
        puts "$expect_out(0,string)"
        puts "}"
    }
}

# Ending script execution
expect eof
