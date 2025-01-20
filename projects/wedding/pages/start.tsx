import Countdown from "@/app/Countdown"
import React from "react"
import styles from "./start.module.css"

export default function Start() {
    return (
            <>
                <div>
                    <h1>Kate &amp; Amrik</h1>
                    <h2>are getting married!</h2>
                </div>
                <div>
                    <Countdown target={new Date("2025-02-02")} />
                </div>
                <div>
                    <a href="#location">Scroll down for more info</a>
                    <p>We will update this website with more info as well as some fun bits closer to the date.</p>
                </div>
                <div className={styles.btnContainer}>
                    <a href="#rsvp" className={styles.btn}>Don&apos;t forget to RSVP.</a>
                </div>
            </>
           )
}
