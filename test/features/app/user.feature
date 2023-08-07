Feature: User administration
  Scenario: Register a new account
    Given I register a new account with username admin and password admin
    Then The registration succeeded
    When I log in to the app using username admin and password admin
    Then The logging is succeeded
    When I log in to the app using username admin and password wrongpassword
    Then The logging is failed
    When I register a new account with username admin and password admin
    Then The registration failed