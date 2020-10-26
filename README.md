# Kenshi Strength Calculator

Allows a user to calculate the minimum Strength level needed to wield certain weapons in the popular RPG [Kenshi](https://lofigames.com/).

See the [community wiki](https://kenshi.fandom.com/wiki/Weapons) for weapon statistics and test data.

## Running the Calculator

1. Clone this repo
1. Navigate into the repo folder
    ```bash
    $ cd KenshiWeaponWeightCalc
    ```
1. Run the command:
    ```go
    $ go run kenshi_strength_calc.go
    ```

## Testing

- Run the following command in the repo folder:
    ```go
    $ go test
    ```

## Development Roadmap

- Add GUI
- Add Weapon objects with associated stats pre-loaded