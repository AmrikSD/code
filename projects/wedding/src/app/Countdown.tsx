import { useEffect, useState } from "react";
import styles from "./countdown.module.css"

type CountdownProps = {
    target: Date
}

type Remaining = {
    days: number,
    hours: number,
    mins: number,
    secs: number
}

const getRemaining = (target: Date): Remaining => {
    const now = new Date()
    const daysUntilTarget = Math.ceil(Math.abs(now.getTime() - target.getTime()) / (1000 * 3600 * 24))
    return {
        days: (daysUntilTarget > 1) ? daysUntilTarget-1 : 0,
        hours: 23 - now.getHours(),
        mins: 59 - now.getMinutes(),
        secs: 59 - now.getSeconds()
    }

}

export default function Countdown(props: CountdownProps){
    const [remaining, setRemaining] = useState<Remaining>(getRemaining(props.target));


    useEffect(() => {
      const interval = setInterval(() => {
          setRemaining(getRemaining(props.target))
      }, 1000);
      return () => {
        clearInterval(interval);
      };
    }, [props.target]);

    return (
        <div className={styles.countdown}>
            <h3>{props.target.toLocaleDateString()}</h3>
            <ul>
                <li>{remaining.days} Days</li>
                <li>{remaining.hours} Hours</li>
                <li>{remaining.mins} Mins</li>
                <li>{remaining.secs} Secs</li>
            </ul>
        </div>
    )
}
