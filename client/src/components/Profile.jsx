import Navbar from "./Navbar";
import ChangePassword from "./ChangePassword";
import NavigateDetailProperty from "./NavbarDetailProperty";
import { Button, Container, Row, Col } from "react-bootstrap";
import React, { useContext, useEffect, useState } from "react";

import Name from "../assets/images/IP.png";
import Email from "../assets/images/Email.png";
import Phone from "../assets/images/Phone.png";
import Place from "../assets/images/Tendant.png";
import Image from "../assets/images/img.jpg";
import Ad from "../assets/images/Lock.jpg";
import G from "../assets/images/Gender.png";
import Password from "../assets/images/Pass.png";
import { UserContext } from "../context/userContext";

export default function Profile(props) {
  useEffect(() => {
    document.body.style.background = "rgba(196, 196, 196, 0.25)";
  });

  const [state, useState] = useContext(UserContext);
  const [show, setShow] = useState(false);
  console.log(state.user);
  const [modalShow, setModalShow] = useState(false);

  return (
    <div className="" style={{ marginTop: "9rem" }}>
      {/* <HomeLogin /> */}
      <Navbar />
      <Container className="">
        <Row className="justify-content-between bg-white py-3">
          <Col className="d-flex flex-column gap-3" sm={4}>
            <h3 className="fw-bold ">Personal Info</h3>
            <div className="d-flex align-items-center gap-3">
              <div>
                <img width={40} src={Name} alt="" />
              </div>
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{state.user?.fullname}</span>
                <span className="fs14 text-secondary">Full Name</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Email} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{state.user?.email}</span>
                <span className="fs14 text-secondary">Email</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Password} alt="" />
              <div className="d-flex align-items-center gap-3">
                <div className="d-flex flex-column">
                  <Button
                    onClick={() => setModalShow(true)}
                    className="btn btn-dark bg-white text-primary fw-bold p-0 m-0 border-0"
                  >
                    Change Password
                  </Button>
                  <ChangePassword
                    show={modalShow}
                    onHide={() => setModalShow(false)}
                  />
                  <span className="p-0 m-0">Password</span>
                </div>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={39} src={Place} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">
                  {state.user?.listAsRole}
                </span>
                <span className="fs14 text-secondary">Status</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={39} src={G} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{state.user?.gender}</span>
                <span className="fs14 text-secondary">Gender</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={40} src={Phone} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{state.user?.phone}</span>
                <span className="fs14 text-secondary">Mobile phone</span>
              </div>
            </div>
            <div className="d-flex align-items-center gap-3">
              <img width={38} src={Ad} alt="" />
              <div className="d-flex flex-column">
                <span className="p-0 m-0 fw-bold">{state.user?.address}</span>
                <span className="fs14 text-secondary">Address</span>
              </div>
            </div>
          </Col>
          <Col className=" d-flex flex-column gap-3 px-3 me-2" sm={4}>
            <div>
              <img className="w-100 rounded pt-4" src={Image} alt="" />
            </div>
            <div>
              <Button className="w-100" onClick={() => setShow(true)}>
                Change Foto Profile
              </Button>
            </div>
          </Col>
        </Row>
      </Container>
    </div>
  );
}
