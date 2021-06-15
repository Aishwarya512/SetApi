import io.gatling.core.Predef._
import io.gatling.http.Predef._

class SetApiSimulation extends Simulation {

  val httpProtocol = http.baseUrl("http://ec2-18-144-33-240.us-west-1.compute.amazonaws.com:3000")

    val scn = scenario("All operations in one scenario")
    .exec(
        http("My Request")
        .get("/addItem/10")
        .check(status.is(400))
    )
    .exec(
        http("My Request")
        .get("/removeItem/294")
        .check(status.is(400))
    )
    .exec(
        http("My Request")
        .get("/hasItem/10")
        .check(status.is(200))
    )
    setUp(scn.inject(atOnceUsers(1000)).protocols(httpProtocol))
}