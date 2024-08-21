#!/bin/bash

# Function to check if the database is running
check_database() {
    # Add your logic to check if the database is running
    # For example, check if the database process is running or if a specific port is open
    # If the database is running, return 0 (success), otherwise return 1 (failure)
    # Example:
    if curl localhost:8200 -s >/dev/null; then
        return 0
    else
        return 1
    fi
    return 0  # Placeholder for demonstration
}

# Function to wait for the database to start
wait_for_database() {
    echo "Waiting for the database to start..."
    while ! check_database; do
        sleep 1
    done
    echo "Database is now running."
}


# Main function
main() {
    # Setup test env using simulator and docker container
    gcloud emulators firestore start --host-port=localhost:8200 -q &
    emulator_pid=$!
    

    # Terminate the Firestore emulator background process
    kill $emulator_pid
}

main