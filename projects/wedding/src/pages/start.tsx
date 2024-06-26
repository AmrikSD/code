import Countdown from "@/app/Countdown"
import React from "react"
import styles from "./start.module.css"

export default function Start() {
    return (
            <>
                <div>
                    <h1>Kate &amp; Amrik</h1>
                    <h2>Are getting married</h2>
                </div>
                <Countdown target={new Date("2025-02-02")} />
                <div className={styles.btnContainer}>
                    <a href="#rsvp" className={styles.btn}>Don&apos;t forget to RSVP.</a>
                </div>
                <div>
                    <p>Scroll down for some <a href="#info">more info</a>!</p>
                    <p><a href="#location">location</a></p>
                    <p><a href="#faq">faq</a></p>
                </div>
            </>
           )
}
