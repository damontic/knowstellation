# Knowstellation

## What is it?
Knowledge Management System (KMS) that allows to track the knowledge database for a set of people. The principal characteristic of the knowstellation is not to be intrusive.

## Definitions
- **star**: specific area of knowledge consisting of multiple caracteristics associated with some value.
    - qa role:
        - validate user stories
        - automatize tests
    - Project A:
        - How to deploy?
        - How to configure tool x?

- **star class**: there are different classes of stars:
    - project stars
    - role stars
    - techniques stars
    - tools stars
    - platform stars
    - languages and frameworks stars

- **star creator**: role assigned to users that can create stars.

- **star visitor**: role assigned to users that have visited stars.

- **star path**: set of stars a user has visited.

- **star evolution**: track of the caracteristics associated to each caracteristics in time. A photo of a star in time. By default, all the star caracteristics decrease in time given some speed.

- **stardust**: element that changes some characteristics of a star.

- **star distance**: measue of how a star is related to another.

## Components
- **knowstellationd**: The star system database and api deamon.

- **knowstellation**: cli that consumes the api.

## How is each star value updated?
- **incoming stardust**: any application that allows plugins can be extended to generate stardust. (slack, jenkins)

- **stardust generators**: middleware applications that read info from other universes and generates stardust. (bash history commands)

- **star agents**: stardust generator agent program installed in each computer that knows how many times a star visitor executes a program and generates stardust.

