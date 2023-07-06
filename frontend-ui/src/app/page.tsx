'use client'

import {SubmitHandler, useForm} from "react-hook-form";
// @ts-ignore
import {saveAs} from "file-saver"

type Inputs = {
    template: string,
    level: string,
    startDate: Date,
    endDate: Date,
    studentName: string,
    marks: number
}

function composeCertDescription(data: Inputs): string {
    if (data.template === "english") {
        return `For his active and invaluable participation during the ${data.level} of the Essential English Course at Omniscience School from ${data.startDate} to ${data.endDate}, and finished the level with ${data.marks} marks.`
    } else if (data.template === "portuguese") {
        return `Por sua participação activa e inestimável ao Curso de Fundamentos de Informática na Omniscience School de ${data.startDate} a ${data.endDate}, tendo terminado o curso com ${data.marks} valores.`
    } else {
        return ""
    }
}

export default function Home() {
    const {register, handleSubmit, formState: {errors}} = useForm<Inputs>()
    const onSubmit: SubmitHandler<Inputs> = function (data) {

        fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/certificates`, {
            method: 'POST',
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                template: data.template,
                studentName: data.studentName,
                description: composeCertDescription(data)
            })
        }).then(res => res.blob()).then(blob => saveAs(blob, `${data.studentName} - Certificate.pdf`.replace(' ',''))).catch(err => {
            console.log(err)
        })
    }

    return (
        <main className="col-md-6 offset-md-3 mt-md-5">
            <h2>Generate a certificate</h2>
            <p className="mb-4"><small>Fill the form bellow with data that will be used to generate the certificate of
                participation.</small></p>

            <form onSubmit={handleSubmit(onSubmit)}>
                <div className="mb-3">
                    <label htmlFor="template" className="form-label">Course</label>
                    <select required defaultValue="english" className="form-select"
                            id="template" {...register('template')}>
                        <option value="english">Essential English Course</option>
                        <option value="portuguese">Computing Fundamentals</option>
                    </select>
                </div>
                <div className="row mb-3">
                    <div className="col">
                        <label htmlFor="level" className="form-label">Level</label>
                        <select defaultValue="Level 1" className="form-select" id="level" {...register('level')}>
                            <option>Level 1</option>
                            <option>Level 2</option>
                            <option>Level 3</option>
                            <option>Level 4</option>
                        </select>
                    </div>
                    <div className="col">
                        <label htmlFor="startDate" className="form-label">Class start date</label>
                        <input required type="date" className="form-control" id="startDate" {...register('startDate')}/>
                    </div>
                    <div className="col">
                        <label htmlFor="endDate" className="form-label">Class end date</label>
                        <input required type="date" className="form-control" id="endDate" {...register('endDate')}/>
                    </div>
                </div>
                <div className="row mb-3">
                    <div className="col-8">
                        <label htmlFor="studentName" className="form-label">Student name</label>
                        <input required type="text" className="form-control" id="studentName"
                               aria-describedby="studentNameHelp" {...register('studentName')}/>
                        <div id="studentNameHelp" className="form-text">This name is of the student being awarded the
                            certificate.
                        </div>
                    </div>
                    <div className="col-4">
                        <label htmlFor="marks" className="form-label">Marks</label>
                        <input required type="number" className="form-control" id="marks" {...register('marks')}/>
                    </div>
                </div>
                <button type="submit" className="btn btn-primary">Submit</button>
            </form>

        </main>
    )
}
