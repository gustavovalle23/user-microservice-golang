Feature: User login

  Scenario: Successful login
    Given a user with username "john" and password "secret"
    When the user logs in with valid credentials
    Then the user should be redirected to the dashboard

  Scenario: Unsuccessful login
    Given a user with username "john" and password "wrongpassword"
    When the user logs in with invalid credentials
    Then the user should see an error message
