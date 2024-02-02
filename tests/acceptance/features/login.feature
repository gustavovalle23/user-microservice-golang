Feature: User login

  Scenario: Successful login
    Given a user with username "gustavo" and password "valle"
    When the user logs in
    Then the user should be logged in successfully

  Scenario: Unsuccessful login
    Given a user with username "gustavo" and password "vale"
    When the user logs in
    Then the user should see an error message
