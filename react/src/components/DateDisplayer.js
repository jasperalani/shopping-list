import React from 'react'

class DateDisplayer extends React.Component {

  constructor (props) {
    super(props)

    this.state = {
      days: [
        'Sun',
        'Mon',
        'Tue',
        'Wed',
        'Thu',
        'Fri',
        'Sat',
      ],
      months: [
        'January',
        'February',
        'March',
        'April',
        'May',
        'June',
        'July',
        'August',
        'September',
        'October',
        'November',
        'December',
      ],
    }
  }

  nth (date) {
    if (date > 3 && date < 21) return 'th';
    switch (date % 10) {
      case 1:  return "st";
      case 2:  return "nd";
      case 3:  return "rd";
      default: return "th";
    }
  }

  componentDidMount () {
    const dateObject = new Date(this.props.date)
    const newDateObject = new Date()
    const objectYear = dateObject.getFullYear()
    const thisYear = newDateObject.getFullYear()
    const year = objectYear === thisYear ? '' : objectYear
    const date = dateObject.getDate() + this.nth(dateObject.getDate())
    const dayName = this.state.days[dateObject.getDay()]
    const monthName = this.state.months[dateObject.getMonth()]
    const formatted = `${dayName}, ${date} ${monthName} ${year}`
    this.setState({ formatted: formatted })
  }

  render () {
    return (
      <span className={'date'}>{this.state.formatted}</span>
    )
  }

}

export default DateDisplayer