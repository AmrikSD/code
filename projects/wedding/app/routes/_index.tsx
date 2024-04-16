import type { ActionFunction, ActionFunctionArgs, LoaderFunction } from "@remix-run/cloudflare";
import { json } from "@remix-run/cloudflare";
import CountDown from "./Countdown";

interface Env {
  MY_DB: D1Database;
}

type rsvp = {
    rsvp_id:number,
    first_name:number,
    last_name:number,
    plus_one:number,
    plus_one_first_name:number,
    plus_one_last_name:number,
    coming_to_thailand:number,
    coming_to_uk:number,
    diet:number
}

export const loader: LoaderFunction = async ({ context })  => {
  const env = context.cloudflare.env as Env
  const { results } = await env.MY_DB.prepare("SELECT * FROM rsvp ORDER BY rowid desc LIMIT 1").all();
  return json(results);
};

export const action: ActionFunction = async ({request, context} : ActionFunctionArgs) => {
    const env = context.cloudflare.env

    const formData = await request.formData()
    const first_name = "first_name"
    const last_name = "last_name"
    const plus_one = 1
    const plus_one_first_name = "plus_one_first_name"
    const plus_one_last_name = "plus_one_last_name"
    const coming_to_thailand = 2
    const coming_to_uk = 2
    const diet = "diet"

    console.log(formData)

    return await env.MY_DB.prepare(`
        INSERT INTO rsvp (first_name, last_name, plus_one, plus_one_first_name, plus_one_last_name, coming_to_thailand, coming_to_uk, diet)
        VALUES (?1, ?2, ?3, ?4, ?5, ?6, ?7, ?8);
    `).bind(
        first_name,
        last_name,
        plus_one,
        plus_one_first_name,
        plus_one_last_name,
        coming_to_thailand,
        coming_to_uk,
        diet,
    ).run().then((resp) => {
        return resp
    })

}

export default function Index() {


  return (
    <>
        <section>
            <div className="container">
                <h1>We&apos;re getting married!</h1>
                <h1>You&apos;re invited!</h1>
                <p>We (<u className="kate">Kate</u> &amp; Amrik) would love to invite you to our wedding!</p>
                <p>For more info, just scroll down!</p>
                <p>To <b>RSVP</b> just scroll down, or <b>TEXT</b> us!</p>
            </div>
        </section>
        <section>
            <div className="container">
                <h1>Hua Hin, Thailand</h1>
                <CountDown target={new Date("2025-02-02")}/>
                <p>yes, a destination wedding :)</p>
                <p>Don&apos;t worry, there will be another celebration in the UK</p>
            </div>
        </section>
        <section>
            <div className="container">
                <button type="submit">RSVP!</button>
            </div>
        </section>
        <section>
            <div className="container">
                <p>Website created with ❤️ and 🫖 by <u className="amrik">Amrik</u></p>
            </div>
        </section>
    </>
  );
}
