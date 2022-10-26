# 1. Context
I am a busy father. When we go to family vacation we always have troubles with packing because none of us thinks about weather. We just pack what we like to wear. Therefore, we would appreciate a bit of help.

# 2. Task overview
Implement a simple service for quering for packing hints. Use wheater API of your choice (You may also provide some kind of mock-service yourself).

## 2.1 Theme
Make family packing easier and faster.

## 2.2 Feature
Possibility to provide a specific date or date range plus location (latitude and longitude is fine, city name is perfect) and retrieve recommended clothes list based on the daily weather.

## 2.3 User stories
* As a busy father I want to receive 1 tshirt & 1 pair of socks and 1 panties per day recommendation for each day in the selected date range and location, so I can wear clean clothes each day.
* As a busy father I want to receive 1 jacket recommendation if any day in the date range has temperature below 11 C°, so I won't freeze out there.
* As a busy father I want to receive 1 umbrella recommandation if any day in the date range will be rainy, so I won't get wet.
* As a busy father I want to receive 1 shorts for each 3-days in the selected date range for which the temperature in given location is higher than 21 C° recommendation, so I can wear clean shorts.
(Example: 7 days in which 6 days have temperature higher than 21 C° = 2 shorts)
* As a busy father I want to receive 1 trousers and 1 sweater recommendation for each 3-days for which the temperature is between 11 C° and 21 C°, so I can wear clean trousers and sweater.
(Example: 7 days in which 6 days have temperature 12 C° = 2 trousers, 2 swaeters)


### 2.3.1 Examples
* INPUT: 03.03.2021 - 09.03.2021, Warszawa
    * Wheater:
        * 03.03 - 10 C°
        * 04.03 - 13 C°
        * 05.03 - 14 C°
        * 06.03 - 12 C°
        * 07.03 - 15 C°
        * 08.03 - 14 C°
        * 09.03 - 17 C°
    * Output: [1 Jacket, 7 Tshirts, 7 pair of socks, 7 panties, 2 Sweaters, 2 Trousers]
* INPUT: 04.04.2021 - 04.04.2021, Sosnowiec
    * Wheater:
        * 04.01 - 22 C°
    * Output: [1 Tshirts]
* INPUT: 04.05.2021 - 07.05.2021, Warszawa
    * Wheater:
        * 04.03 - 15 C°
        * 05.03 - 22 C°
        * 06.03 - 22 C° Rainy
        * 07.03 - 27 C°
    * Output: [4 Tshirts, 1 Sweaters, 1 Trousers, 1 Umbrella]

# 3. Technical Requirements:
- HTTP API

# 4. Additional Notes:
- Use common sense if something wasn't specified