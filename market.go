package nyse_time

import (
	"time"
)

/**SECTION - Main Function
 * * @desc The following function lets you know if the NYSE market is open.
 * @note These functions are exported and can be used by other packages.
 */

/**ANCHOR - IsMarketOpen
 * @return bool
 * * @desc Returns true if the market is open on the given date, false otherwise.
 * @note The market is open from 9:30am to 4:00pm EST/EDT on weekdays.
 * @note The market is closed on weekends and holidays.
 * @note The market closes at 1:00pm EST/EDT on the day before Thanksgiving.
 * @note The market closes at 1:00pm EST/EDT on Christmas Eve.
 * @note The market closes at 1:00pm EST/EDT on the day before Independence Day.
 */
func IsMarketOpen() bool {
	currentTime := time.Now().UTC()
	return !IsWeekend(currentTime) && !IsHoliday(currentTime) && IsNormalHours(currentTime)
}

/**ANCHOR - IsMarketOpenCustom
 * @param t time.Time
 * @return bool
 * * @desc Returns true if the market is open on the given date, false otherwise.
 * @note The market is open from 9:30am to 4:00pm EST/EDT on weekdays.
 * @note The market is closed on weekends and holidays.
 * @note The market closes at 1:00pm EST/EDT on the day before Thanksgiving.
 * @note The market closes at 1:00pm EST/EDT on Christmas Eve.
 * @note The market closes at 1:00pm EST/EDT on the day before Independence Day.
 */
func IsMarketOpenCustom(t time.Time) bool {
	return !IsWeekend(t) && !IsHoliday(t) && IsNormalHours(t)
}

//!SECTION
