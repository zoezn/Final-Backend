<h1>Especialización en Back End III</h1>
    <h2>Taller de código: Desafío final</h2>
    <p>
      A continuación se plantea un desafío integrador que nos permitirá evaluar
      todos los temas que hemos visto en la cursada.
    </p>
    <h3>Sistema de reserva de turnos</h3>
    <p>
      Se desea implementar una API que permita administrar la reserva de turnos
      para una clínica odontológica. Esta debe cumplir con los siguientes
      requerimientos:
    </p>
    <ul>
      <li type="disc">
        <div>
          <p>
            <b>  Administración de datos de odontólogos:</b>
            listar, agregar, modificar y
            eliminar odontólogos. Registrar apellido, nombre y matrícula de los
            mismos. Se desea el desarrollo de un 
            <b>CRUD </b> para la entidad Dentista.
          </p>
          <ul type="circle">
            <li>POST: agregar dentista.</li>
            <li>GET: traer dentista por ID.</li>
            <li>PUT: actualizar dentista.</li>
            <li>PATCH: actualizar un dentista por alguno de sus campos.</li>
            <li>DELETE: eliminar el dentista.</li>
          </ul>
        </div>
      </li>
      <li type="disc">
        <div>
          <p>
            <b> Administración de datos de los pacientes: </b>
            listar, agregar, modificar y eliminar pacientes. De cada uno se
            almacenan: nombre, apellido, domicilio, DNI y fecha de alta. Se
            desea el desarrollo de un
            <b>CRUD </b> para la entidad Paciente.
          </p>
          <ul type="circle">
            <li>POST: agregar paciente.</li>
            <li>GET: traer paciente por ID.</li>
            <li>PUT: actualizar paciente.</li>
            <li>PATCH: actualizar un paciente por alguno de sus campos.</li>
            <li>DELETE: eliminar al paciente.</li>
          </ul>
        </div>
      </li>
      <li type="disc">
        <div>
          <p>
            <b> Registrar turno:</b>
            se tiene que poder permitir asignar a un paciente un turno con un
            odontólogo a una determinada fecha y hora. Al turno se le debe poder
            agregar una descripción. Se desea el desarrollo de un
            <b>CRUD </b> para la entidad Turno.
          </p>
          <ul type="circle">
            <li>POST: agregar turno.</li>
            <li>GET: traer turno por ID.</li>
            <li>PUT: actualizar turno.</li>
            <li>PATCH: actualizar un turno por alguno de sus campos.</li>
            <li>DELETE: eliminar turno.</li>
            <li>
              POST: agregar turno por DNI del paciente y matrícula del dentista.
            </li>
            <li>
              GET: traer turno por DNI del paciente. Debe traer el detalle del
              turno (Fecha-Hora, descripción, Paciente y Dentista) y el dni
              deberá ser recibido por QueryParams.
            </li>
          </ul>
        </div>
      </li>
      <li type="disc">
        <p>
          <b>Seguridad mediante middleware:</b> se tiene que proveer cierta
          seguridad al momento de realizar <b> POST, PUT, PATCH y DELETE</b>.
          Esta seguridad mediante autenticación deberá estar implementada
          mediante un middleware.
        </p>
      </li>
      <li type="disc">
        <p>
          <b> Documentación de la API: </b>
          se debe proveer de la pertinente documentación de la API mediante la
          implementación de <b> Swagger</b>.
        </p>
      </li>
    </ul>
      <h3>Requerimientos técnicos</h3>
    <p>La aplicación debe ser desarrollada en diseño orientado a paquetes:</p>
    <ul type="disc">
      <li><b> Capa/dominio de entidades de negocio. </b></li>
      <li>
        <b> Capa de acceso a datos (base de datos): </b>es la base de datos de
        nuestro sistema. Podrás utilizar cualquier base de datos relacional
        modelado a través de un modelo entidad-relación, como H2 o MySQL, o no
        relacional, como MongoDB.
      </li>
      <li><b> Capa/dominio service. </b></li>
      <li><b> Capa/dominio handler. </b></li>
    </ul>